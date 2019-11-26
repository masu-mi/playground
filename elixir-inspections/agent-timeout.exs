Agent.start_link(fn -> 0 end, name: Sample)
|> IO.inspect()

defmodule ExampleCode do
  def agent_timeout(update_timeout_msec, sleep_time \\ 1000) do
    try do
      result = Agent.update(
        Sample,
        fn state ->
          :timer.sleep(sleep_time)
          state + 1
        end,
        update_timeout_msec
      )
      {:ok, update_timeout_msec, sleep_time, result}
    catch
      :exit, e ->
        {:exit, update_timeout_msec, sleep_time, e}
    end
  end
end
