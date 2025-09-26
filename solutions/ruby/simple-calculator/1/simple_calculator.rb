class SimpleCalculator
  class UnsupportedOperation < StandardError; end

  ALLOWED_OPERATIONS = ['+', '/', '*'].freeze

  def self.calculate(first_operand, second_operand, operation)
    unless ALLOWED_OPERATIONS.include?(operation)
      raise UnsupportedOperation.new
    end

    unless first_operand.is_a?(Numeric) && second_operand.is_a?(Numeric)
      raise ArgumentError.new
    end

    if operation == '/' && second_operand == 0
      return "Division by zero is not allowed."
    end

    process_str = "#{first_operand} #{operation} #{second_operand}"
    "#{process_str} = #{eval(process_str)}"
  end
end