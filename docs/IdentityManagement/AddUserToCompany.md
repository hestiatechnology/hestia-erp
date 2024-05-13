# Add a user to a company
```mermaid
graph TD
    A[Start] --> G{Is there a user with the given email}
    G --> |No| H[Return 'User not found']


    G --> |Yes| I{Is the user already in the company}
    I --> |Yes| J[Return 'User already in the company']
    I --> |No| K{Is employeeId provided}
    K --> |employeeId is not provided| L[Associate user with the company without employeeId]
    K --> |employeeId is provided| M{Check if employeeId is already in use}
    M --> |employeeId is in use| N[Return 'Employee ID already in use']
    M --> |employeeId is not in use| O[Associate user with the company with employeeId]
    L --> Q[End]
    O --> Q[End]
```