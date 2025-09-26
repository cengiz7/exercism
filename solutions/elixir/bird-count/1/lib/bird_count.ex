defmodule BirdCount do
  def today([]), do: nil

  def today(list) do
    List.first(list)
  end

  def increment_day_count(list) do
    case today(list) do
      nil -> [1]
      _ -> List.update_at(list, 0, &(&1 + 1))
    end
  end

  def has_day_without_birds?(list) do
    0 in list
  end

  def total(list) do
    Enum.reduce(list, 0, &(&1 + &2))
  end

  def busy_days(list) do
    Enum.count(list, &(&1 > 4))
  end
end
