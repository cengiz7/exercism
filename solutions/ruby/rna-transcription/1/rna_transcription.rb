module Complement
  COMPLEMENTS = {
    'G' => 'C',
    'C' => 'G',
    'T' => 'A',
    'A' => 'U'
  }
  
  def self.of_dna(dna)
    dna.upcase.chars.map {|c| COMPLEMENTS[c] }.join
  end
end