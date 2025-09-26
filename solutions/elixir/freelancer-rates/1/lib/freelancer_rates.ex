defmodule FreelancerRates do
  def daily_rate(hourly_rate) do
    8.0 * hourly_rate
  end

  def apply_discount(before_discount, discount) do
    before_discount - (before_discount * discount / 100)
  end

  def monthly_rate(hourly_rate, discount) do
    hourly_rate
    |> daily_rate
    |> Kernel.*(22)
    |> apply_discount(discount)
    |> ceil
  end

  def days_in_budget(budget, hourly_rate, discount) do
    budget / (hourly_rate
    |> daily_rate()
    |> apply_discount(discount))
    |> Float.floor(1)
  end
end
