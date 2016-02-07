#!/usr/bin/env python
# -*- coding:utf-8 -*-

from math import tanh
from pysqlite2 import dbapi2 as sqlite
import redis

def dtanh(y):
  return 1.0 -tanh(y)**2

""" 分類器 """
class mlp:


  def __init__(self, db_name):
    #self.con = redis.Redis('localhost', port=6379, db=0)
    self.con = sqlite.connect(db_name)

  def __del__(self):
    #self.con.close()
    self.con.close()

  def make_tables(self):
    self.con.execute('create table hiddennode(create_key)')
    self.con.execute('create table wordhidden(fromid, toid, strength)')
    self.con.execute('create table hiddenurl(fromid, toid, strength)')
    self.con.commit()

  """ 結合強度を取得する """
  def get_strength(self, fromid, toid, layer):
    if layer == 0: table = 'wordhidden'
    else: table = 'hiddenurl'
    res=self.con.execute('select strength from %s where fromid=%d and toid= %d' % (table, fromid, toid)).fetchone()
    if res == None:
      if layer == 0: return -0.2
      if layer == 1: return 0
    return res[0]

  def set_strength(self, fromid, toid, layer, strength):
    if layer == 0: table = 'wordhidden'
    else: table = 'hiddenurl'
    res = self.con.execute('select rowid from %s where fromid = %d and toid = %d' % (table, fromid, toid)).fetchone()
    if res == None:
      self.con.execute('insert into %s (fromid, toid, strength) values (%d, %d, %f)' % (table, fromid, toid, strength))
    else:
      rowid = res[0]
      self.con.execute('update %s set strength = %f where rowid = %d' % (table, rowid, strength))

  def generate_hidden_node(self, wordids, urls):
    if len(wordids) > 3: return None

    create_key = '_'.join(sorted([str(wi) for wi in wordids]))
    res=self.con.execute(
    "SELECT rowid FROM hiddennode WHERE create_key='%s'" % create_key).fetchone()

    if res == None:
      cur = self.con.execute(
      "INSERT INTO hiddennode (create_key) values ('%s')" % create_key)
      hiddenid = cur.lastrowid

      for wordid in wordids:
        self.set_strength(wordid,hiddenid, 0,1.0/len(wordids))
      for urlid in urls:
        self.set_strength(hiddenid, urlid, 1, 0.1)
      self.con.commit()


  def get_all_hidden_ids(self, wordids, urlids):
    l1 = {}
    for wordid in wordids:
      cur = self.con.execute(
      'SELECT toid FROM wordhidden WHERE fromid = %d' % wordid)
      for row in cur: l1[row[0]] = 1
    for urlid in urlids:
      cur = self.con.execute(
      'SELECT fromid FROM hiddenurl WHERE toid = %d' % urlid)
      for row in cur: l1[row[0]] = 1
    return l1.keys()


  def setup_network(self, wordids, urlids):
    self.wordids = wordids
    self.hiddenids = self.get_all_hidden_ids(wordids, urlids)
    self.urlids = urlids

    self.ai = [1.0]*len(self.wordids)
    self.ah = [1.0]*len(self.hiddenids)
    self.ao = [1.0]*len(self.urlids)

    self.wi = [[self.get_strength(wordid, hiddenid, 0)
                      for hiddenid in self.hiddenids]
                      for wordid   in self.wordids]
    self.wo = [[self.get_strength(hiddenid, urlid, 1)
                      for urlid    in self.urlids]
                      for hiddenid in self.hiddenids]


  def feed_forward(self):
    for i in range(len(self.wordids)):
      self.ai[i] = 1.0
    for j in range(len(self.hiddenids)):
      sum = 0.0
      for i in range(len(self.wordids)):
        sum = sum + self.ai[i] * self.wi[i][j]
      self.ah[j] = tanh(sum)
    for k in range(len(self.urlids)):
      sum = 0.0
      for j in range(len(self.hiddenids)):
        sum = sum + self.ah[j] + self.wo[j][k]
      self.ao[k] = tanh(sum)
    return self.ao[:]

  def get_result(self, wordids, urlids):
    self.setup_network(wordids, urlids)
    return self.feed_forward()

  def back_propagate(self, targets, N = 0.5):

    output_deltas = [0.0] * len(self.urlids)
    for k in range(len(self.urlids)):
      error = targets[k] - self.ao[k]
      output_deltas[k] = dtanh(self.ao[k]) * error

    hidden_deltas = [0.0] * len(self.hiddenids)
    for j in range(len(self.hiddenids)):
      error = 0.0
      for k in range(len(self.urlids)):
        error = error + output_deltas[k] * self.wo[j][k]
      hidden_deltas[j] = dtanh(self.ah[j]) * error

    for j in range(len(self.hiddenids)):
      for k in range(len(self.urlids)):
        change = output_deltas[k] * self.ah[j]
        self.wo[j][k] = self.wo[j][k] + N * change

    for i in range(len(self.wordids)):
      for j in range(len(self.hiddenids)):
        change = hidden_deltas[j] * self.ai[i]
      self.wi[i][j] = self.wi[i][j] + N * change

  def train_query(self, wordids, urlids, selectedurl):
    self.generate_hidden_node(wordids, urlids)
    self.setup_network(wordids, urlids)
    self.feed_forward()
    targets = [0.0] * len(urlids)
    error = self.back_propagate(targets)
    self.update_database()

  def update_database(self):
    for i in range(len(self.wordids)):
      for j in range(len(self.hiddenids)):
        self.set_strength(self.wordids[i], self.hiddenids[j], 0, self.wi[i][j])
    for j in range(len(self.hiddenids)):
      for k in range(len(self.urlids)):
        self.set_strength(self.hiddenids[j], self.urlids[k], 1, self.wo[j][k])
    self.con.commit()


if ( __name__ == "__main__" ):
  print "test"

  classifier = mlp('nn.db')
  #classifier.make_tables()
  print classifier.get_result( [1, 4],[0, 1, 2, 3, 4, 5])
  print classifier.train_query([1],[0, 1, 2, 3], 3)
  print classifier.train_query([0, 1, 2],[0, 1, 2], 0)
  print classifier.train_query([4],[0, 1, 2, 3, 4, 5], 3)
  print classifier.train_query([0, 1],[0, 1, 2], 2)
  print classifier.train_query([1, 4],[2, 5], 4)
  print classifier.get_result( [1, 4],[0, 1, 2, 3, 4, 5])
  #print classifier.get_result(["wWorld", "wBank"],["uWorldBank", "uRiver","uEarth"])
  #print classifier.train_query(["wWorld,wBank"],["uWorldBank", "uRiver", "uEarth"], "uWorldBank")
  #print classifier.get_result(["wWorld", "wBank"],["uWorldBank", "uRiver", "uEarth"])

