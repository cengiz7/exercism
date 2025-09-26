defmodule BirdCount do
  def today([]), do: nil
  def today([f | _]), do: f

  def increment_day_count([]), do: [1]
  def increment_day_count([f | tail]), do: [f+1 | tail]

  def has_day_without_birds?([]), do: false
  def has_day_without_birds?([0 | _]), do: true
  def has_day_without_birds?([_ | t]) do
    has_day_without_birds?(t)
  end


  def total([]), do: 0
  def total([f | l]), do: f + total(l)

  def busy_days([]), do: 0
  def busy_days([f | l]) when f > 4, do: 1 + busy_days(l)
  def busy_days([_ | l]), do: busy_days(l)
end
