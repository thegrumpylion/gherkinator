Feature: learn react default page

  Scenario: see the default page
    When I navigate to "/"
    Then I see a link "Learn React"

  Scenario Outline: eating
    Given there are <start> cucumbers
    When I eat <eat> cucumbers
    Then I should have <left> cucumbers
    Examples:
      | start | eat | left |
      |    12 |   5 |    7 |
      |    20 |   5 |   15 |

  Scenario: login with valid credentials
    Given a user "asdfasd" has been created with the following details:
       | email          | username  | password |
       | user@email.com | user      | password |
    And the user has browsed to the login page
    When the user enters the following details in the login form:
       | email          | username  | password |
       | user@email.com | user      | password |
    And the user logs in
    Then the user should be redirected to the homepage
