CREATE TABLE Employees
(
	ID INTEGER PRIMARY KEY,
	Name TEXT,
	Born TIMESTAMP
);

CREATE TABLE Users
(
	Login TEXT PRIMARY KEY
);

CREATE TABLE Subscriptions
(
	UserLogin TEXT,
	EmployeeID INTEGER,
	FOREIGN KEY (UserLogin) REFERENCES Users (login),
	FOREIGN KEY (EmployeeID) REFERENCES Employees(ID)
);
