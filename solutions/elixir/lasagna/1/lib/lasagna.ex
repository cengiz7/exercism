defmodule Lasagna do
  def expected_minutes_in_oven do
    40
  end

  def remaining_minutes_in_oven(past) do
    40 - past
  end

  def preparation_time_in_minutes(layers) do
    2 * layers
  end

  def total_time_in_minutes(layers, spent) do
    2 * layers + spent
  end

  def alarm do
    "Ding!"
  end
end
