module SavingsAccount

  def self.interest_rate(balance)
    case balance
    when -Float::INFINITY...0 then -3.213
    when 0...1000 then 0.5
    when 1000...5000 then 1.621
    when 5000...Float::INFINITY then 2.475
    end
  end

  def self.annual_balance_update(balance)
    balance + (balance.abs / 100) * interest_rate(balance)
  end

  def self.years_before_desired_balance(current_balance, desired_balance)
    calculate_years_for_balance(current_balance, desired_balance)
  end

  private

  def self.calculate_years_for_balance(current_balance, desired_balance, years=0)
    if current_balance >= desired_balance
      years
    else
      new_balance = annual_balance_update(current_balance)
      calculate_years_for_balance(new_balance, desired_balance, years+1)
    end
  end
end