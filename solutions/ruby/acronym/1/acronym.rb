class Acronym
  def self.abbreviate(phrase)
    phrase.gsub!(/[^a-z]/i, ' ').upcase.split.map{|v| v[0]}.join
  end
end