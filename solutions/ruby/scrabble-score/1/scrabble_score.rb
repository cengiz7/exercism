class Scrabble
  SCORES = {'AEIOULNRST' => 1, 'DG' => 2, 'BCMP' => 3, 'FHVWY' => 4, 'K' => 5, 'JX' => 8, 'QZ' => 10}
  
  def initialize(input)
    @input = input
  end

  def score
    @input.to_s.strip.upcase.chars.map{|ch| SCORES.select{|k, v| k.include? ch}.values.first }.sum
  end

  def self.score(input)
    Scrabble.new(input).score
  end
end