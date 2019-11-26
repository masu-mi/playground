defmodule ExampleCode do

  def example_await(timeout, sleep_time \\ 1000) do
    try do
      result = Task.async(fn ->
        :timer.sleep(sleep_time)
        :ok
      end)
      |> Task.await(timeout)
      {:ok, timeout, sleep_time, result}
    catch
      :exit, e -> {:exit, timeout, sleep_time, e}
    end
  end

  def example_yield(timeout, sleep_time \\ 1000) do
    try do
      result = Task.async(fn ->
        :timer.sleep(sleep_time)
        :ok
      end)
      |> Task.yield(timeout)
      {:ok, timeout, sleep_time, result}
    catch
      :exit, e -> {:exit, timeout, sleep_time, e}
    end
  end
end

ExampleCode.example_yield(1)
|> IO.inspect()

ExampleCode.example_await(1)
|> IO.inspect()
