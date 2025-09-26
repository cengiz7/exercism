class SpaceAge
  EARTH_YEAR_SECS = 31557600.to_f
  ORBITAL_PERIODS = {
    'earth': 1, 'mercury': 0.2408467,
    'venus': 0.61519726, 'mars': 1.8808158,
    'jupiter': 11.862615, 'saturn': 29.447498,
    'uranus':  84.016846, 'neptune': 164.79132
  }

  def initialize(seconds)
    @earth_year = seconds / EARTH_YEAR_SECS
  end

  ORBITAL_PERIODS.each do |planet, period|
    define_method("on_#{planet}") do
      @earth_year / period
    end
  end

end