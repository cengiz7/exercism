module Raindrops
  def Raindrops.convert(number)
    result = pling(number) + plang(number) + plong(number)
    result.empty? ? number.to_s : result
  end

  private
  def self.pling(number)
    number % 3 == 0 ? 'Pling' : ''
  end

  def self.plang(number)
    number % 5 == 0 ? 'Plang' : ''
  end

  def self.plong(number)
    number % 7 == 0 ? 'Plong' : ''
  end
end