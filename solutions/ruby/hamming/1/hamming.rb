class Hamming
  def Hamming.compute(dna1, dna2)
    raise ArgumentError.new if dna1.length != dna2.length

    difference_count = 0
    dna1.chars.each_with_index do |char, idx|
      if dna2[idx] != char
        difference_count += 1 
      end
    end
    difference_count
  end
end