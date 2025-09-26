defmodule KitchenCalculator do
  Module.register_attribute __MODULE__, :conversion_chart, accumulate: true

  @conversion_chart {:milliliter, 1}
  @conversion_chart {:cup, 240}
  @conversion_chart {:fluid_ounce, 30}
  @conversion_chart {:teaspoon, 5}
  @conversion_chart {:tablespoon, 15}
  
  def get_volume(volume_pair), do: elem(volume_pair, 1)

  def to_milliliter({unit, vol}), do: {:milliliter, vol * get_conversion_volume(unit)}

  def from_milliliter({_, vol}, unit), do: {unit, vol / get_conversion_volume(unit)}

  def convert({from_unit, vol}, to_unit) do
    {from_unit, vol}  
    |> to_milliliter()
    |> from_milliliter(to_unit)
  end

  defp get_conversion_volume(unit), do: @conversion_chart[unit]
end
