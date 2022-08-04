Feature: learn react default page
  Scenario: login with valid credentials
    Given a user "asdfasd" has been created with the following details:
       | email          | username  | password |
       | user@email.com | user      | password |
    And the user has browsed to the login page
    When the user enters the following details in the login form:
       | email          | username  | password |
       | user@email.com | user      | password |
    And the user logs in
    Then the user 99 be redirected to the homepage
    """
    Yo this is a doc string

    is multi line
    yay
    """
