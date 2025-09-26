module Luhn
    def self.valid?(input)
        input.gsub!(/\s/, '')

        return false if input.length < 2 || input.match(/[^0-9]/)

        numbers = []
        input.reverse.chars.each_with_index do |num, i| 
            num = num.to_i
            if i % 2 == 0
                num
            else
                num = num * 2
                num = num > 9 ? num-9 : num
            end
            numbers.push num
        end
        numbers.sum % 10 == 0
    end
end
