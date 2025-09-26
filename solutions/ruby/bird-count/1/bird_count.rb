class BirdCount
  def self.last_week
    [0, 2, 5, 3, 7, 8, 4]
  end

  def initialize(birds_per_day)
    @this_week_counts = birds_per_day
  end

  def yesterday
    @this_week_counts[-2]
  end

  def total
    @this_week_counts.sum
  end

  def busy_days
    visits = @this_week_counts.select { |x| x >= 5 }
    visits.count
  end

  def day_without_birds?
    @this_week_counts.any? 0
  end
end
