class Series
    def initialize(str)
        @str=str
    end

    def slices(length)
        raise ArgumentError.new if @str.length < length
        (0..(@str.length - length)).map {|i| @str[i, length] }
    end
end