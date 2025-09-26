defmodule FreelancerRates do
  @daily_rate  8.0
  @monthly_billable_days  22

  def daily_rate(hourly_rate), do: @daily_rate * hourly_rate 

  def apply_discount(before_discount, discount) do
    before_discount - (before_discount * discount / 100)
  end

  def monthly_rate(hourly_rate, discount) do
    @monthly_billable_days * daily_discounted_rate(hourly_rate, discount)
    |> ceil
  end

  def days_in_budget(budget, hourly_rate, discount) do
    budget / daily_discounted_rate(hourly_rate, discount)
    |> Float.floor(1)
  end

  defp daily_discounted_rate(hourly_rate, discount) do
    hourly_rate
    |> daily_rate()
    |> apply_discount(discount)
  end
end
