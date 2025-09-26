class SumOfMultiples
  def initialize(*multipliers)
    @multipliers = multipliers
    @multipliers.delete(0)
  end

  def to(number)
    @multipliers.map {|m|
      (1..get_devider(number, m)).map{|v| v * m}
    }.flatten.uniq.sum
  end

  private

  def get_devider(number, multiplier)
    if number % multiplier == 0
      number / multiplier - 1
    else
      number / multiplier
    end
  end
end