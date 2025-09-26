module Port
  IDENTIFIER = :PALE

  def self.get_identifier(city)
    city.upcase[0...4].to_sym
  end

  def self.get_terminal(ship_identifier)
    case ship_identifier.upcase[0...3]
    when /^(OIL|GAS)/
      :A
    else
      :B
    end
  end
end
