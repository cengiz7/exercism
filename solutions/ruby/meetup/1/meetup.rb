class Meetup
  require 'date'

  DAYS = { sunday: 0, monday: 1, tuesday: 2, wednesday: 3, thursday: 4, friday: 5, saturday: 6 }.freeze
  ORDERED_INDICES = { first: 0, second: 1, third: 2, fourth: 3, fifth: 4, last: -1}.freeze

  def initialize(month, year)
    @year  = year
    @month = month
    @date  = Date.new(year, month)
  end

  def day(day, indice)
    indice == :teenth ? match_teenth(day) : week_day_by_count(day, indice)
  end

  private

  def match_teenth(day)
    week_day  = DAYS.fetch(day)
    month_day = 13

    while month_day < 20
      new_date = Date.new(@year, @month, month_day)
      return new_date if new_date.wday == week_day
      month_day += 1
    end
  end

  def week_day_by_count(day, indice)
    week_days = (0...7).map { |i| [i,[]] }.to_h
    for d in 1..Date.new(@year, @month, -1).day do
      date = Date.new(@year, @month, d)
      week_days[date.wday] << date.mday
    end
    Date.new(@year, @month, week_days[DAYS[day]][ORDERED_INDICES[indice]] )
  end
end