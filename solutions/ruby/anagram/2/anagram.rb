class Anagram
  def initialize(word)
    @base_word = word.downcase
    @base_word_sorted = @base_word.chars.sort
  end

  def match(words)
    words.select do |word|
      word = word.downcase
      word != @base_word && @base_word_sorted == word.chars.sort
    end
  end
end
