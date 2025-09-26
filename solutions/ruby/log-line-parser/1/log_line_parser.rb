class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    @line.to_s.split(":").last.strip
  end

  def log_level
    @line.to_s.split(":").first.gsub(/[\[\]]/, "").strip.downcase
  end

  def reformat
    "#{message} (#{log_level})"
  end
end
