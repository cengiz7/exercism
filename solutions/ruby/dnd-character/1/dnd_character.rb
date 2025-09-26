class DndCharacter
  ABILITIES =  %i[hitpoints strength dexterity constitution intelligence wisdom charisma]

  attr_reader(*ABILITIES)

  def self.modifier(constitution)
    ((constitution - 10) / 2).floor
  end

  def initialize
    %i[strength dexterity constitution intelligence wisdom charisma].each do |attribute|
      instance_variable_set(
        "@#{attribute}",
        get_max_three_dice_sum(roll_the_dices)
      )
    end
    @hitpoints = 10 + DndCharacter.modifier(@constitution)
  end

  private

  def roll_the_dices
    4.times.map { rand(1..6) }
  end

  def get_max_three_dice_sum(values)
    values.max(3).sum
  end
end
