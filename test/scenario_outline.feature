Feature: scenario outline eating

  Scenario Outline: eating
    Given there are <start> cucumbers
    When I eat <eat> cucumbers
    Then I should have <left> cucumbers
    Examples:
      | start | eat | left |
      |    12 |   5 |    7 |
      |    20 |   5 |   15 |
      |    27 |   7 |   20 |
  
  Scenario Outline: remove letters
    Given there word <word>
    When I remove letter <letter>
    Then I should have word <left>
    Examples:
      | word  | letter | left |
      | hello | e      | hllo |
      | cruel | l      | crue |
      | world | w      | orld |
