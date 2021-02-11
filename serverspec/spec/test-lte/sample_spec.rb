require 'spec_helper'

describe routing_table do
  it {
    should have_entry(
      :destination => "default",
      :interface => "enp0s3",
      :gateway => "10.0.2.2",
    )
  }
end
