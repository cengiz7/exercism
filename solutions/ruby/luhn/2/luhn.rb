module Luhn
    def self.valid?(input)
        input.gsub!(/\s/, '')

        return false if input.length < 2 || input.match(/[^0-9]/)

        input
        .reverse
        .chars
        .map(&:to_i)
        .map.with_index { |num, index| index.even? ? num : num*2 }
        .map {|num| num>9 ? num-9 : num}
        .sum % 10 == 0
    end
end