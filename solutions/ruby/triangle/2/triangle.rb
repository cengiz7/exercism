class Triangle
  def initialize(sides)
    @sides = sides
    @max = @sides.max
  end

  def equilateral?
    check_sides && @sides.uniq.size == 1
  end

  def isosceles?
    check_sides && @sides.uniq.size <= 2
  end

  def scalene?
    check_sides && @sides.uniq.size == 3
  end

  private
  def check_sides
    a,b,c = @sides
    return false if @sides.any? { |side| side <= 0 }
    (a + b <= c) || (a + c <= b) || (b + c <= a) ? false : true
  end

end
