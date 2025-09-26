class Clock
    attr_accessor :hour, :minute

    def initialize(hour: 0, minute: 0)
        @hour, @minute = hour, minute
        set_minute_and_hour
    end

    def to_s
        sprintf("%02d:%02d", @hour, @minute)
    end

    def -(clock)
        @hour -= clock.hour
        @minute -= clock.minute
        set_minute_and_hour
        self
    end

    def +(clock)
        @hour += clock.hour
        @minute += clock.minute
        set_minute_and_hour
        self
    end

    def ==(clock)
        @hour.eql?(clock.hour) && @minute.eql?(clock.minute)
    end


    private

    def set_minute_and_hour
        day_minutes = (@hour * 60 + @minute) % 1440
        @minute = day_minutes % 60
        @hour   = (day_minutes / 60).floor
    end
end