class Anagram
  def initialize(word)
    @word = word.downcase.chars
    @word_char_map = @word.uniq.each_with_object({}) do |char, hash|
      hash[char] = @word.count(char)
    end
  end

  def match(words)
    words.select do |word|
      test_word = word.downcase.chars
      !@word.eql?(test_word) && test_word.uniq.all? { |char| @word_char_map[char] == test_word.count(char) }
    end
  end
end
