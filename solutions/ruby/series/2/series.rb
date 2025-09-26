class Series
    def initialize(str)
        @str=str
    end

    def slices(length)
        raise ArgumentError.new if @str.length < length
        @str.chars.each_cons(length).map(&:join)
    end
end