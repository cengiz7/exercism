class Phrase
  def initialize(phrase)
    @phrase = phrase
  end

  def word_count
    clean_str = -> (str) { str.strip.gsub(/^[\W\_]/, '').gsub(/[\W\_]+$/, '') }
    @phrase.downcase.split(/\s|,|\R/).filter{|e| !e.empty? }.map(&clean_str).each_with_object(Hash.new(0)) {|word, hash| hash[word] += 1}
  end
end