import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

export default function EmployeeList() {
  const [employees, setEmployees] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetch("http://localhost:8080/api/employee", {
      credentials: "include", // если используешь куки
    })
      .then((res) => res.json())
      .then((data) => {
        setEmployees(data);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Ошибка при получении сотрудников:", err);
        setLoading(false);
      });
  }, []);

  if (loading) return <p>Загрузка...</p>;

  return (
    <div>
      <h2>Список сотрудников</h2>

      <table border="1" cellPadding="5" cellSpacing="0">
        <thead>
          <tr>
            <th>Email</th>
            <th>Имя</th>
            <th>Фамилия</th>
            <th>Отчество</th>
            <th>Адрес</th>
            <th>Телефон</th>
            <th>Дата рождения</th>
            <th>Дата приёма</th>
            <th>Должность</th>
            <th>Действия</th>
          </tr>
        </thead>
        <tbody>
          {employees.length > 0 ? (
            employees.map((emp) => (
              <tr key={emp.ID}>
                <td>{emp.Email}</td>
                <td>{emp.FirstName}</td>
                <td>{emp.LastName}</td>
                <td>{emp.MiddleName}</td>
                <td>{emp.Address}</td>
                <td>{emp.PhoneNumber}</td>
                <td>{emp.DateOfBirth ? emp.DateOfBirth.split("T")[0] : ""}</td>
                <td>{emp.DateOfHiring ? emp.DateOfHiring.split("T")[0] : ""}</td>
                <td>{emp.Position?.Name}</td>
                <td>
                  <Link to={`/employee/EditEmployee/${emp.ID}`}>✏️ Редактировать</Link>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan="10">Сотрудников нет</td>
            </tr>
          )}
        </tbody>
      </table>

      <br />
      <Link to="/employee/CreateEmployee">➕ Добавить сотрудника</Link>
    </div>
  );
}
