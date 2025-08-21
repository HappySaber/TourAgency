// src/pages/employee/CreateEmployee.jsx
import { useState } from "react";

export default function CreateEmployee() {
  const [form, setForm] = useState({
    email: "",
    password: "",
    firstName: "",
    lastName: "",
    middleName: "",
    address: "",
    phoneNumber: "",
    dateOfBirth: "",
    dateOfHiring: "",
    position: "",
  });

  const [positions, setPositions] = useState([]);

  // Загрузим должности при монтировании
  useState(() => {
    fetch("/api/positions") // например, эндпоинт для должностей
      .then((res) => res.json())
      .then((data) => setPositions(data))
      .catch(console.error);
  }, []);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const formatToTimestamp = (dateStr) =>
    dateStr ? `${dateStr}T00:00:00Z` : null;

  const handleSubmit = async (e) => {
    e.preventDefault();

    const data = {
      ...form,
      dateOfBirth: formatToTimestamp(form.dateOfBirth),
      dateOfHiring: formatToTimestamp(form.dateOfHiring),
    };

    const response = await fetch("/api/employeeCreate", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(data),
    });

    const result = await response.json();

    if (response.ok) {
      alert(result.success || "Регистрация прошла успешно");
      window.location.href = "/employee"; // перенаправляем на список сотрудников
    } else {
      alert(result.error || "Ошибка регистрации");
    }
  };

  return (
    <div>
      <h2>Создание сотрудника</h2>
      <form onSubmit={handleSubmit}>
        <label>Email:</label><br />
        <input type="email" name="email" value={form.email} onChange={handleChange} required /><br /><br />

        <label>Пароль:</label><br />
        <input type="password" name="password" value={form.password} onChange={handleChange} required minLength={6} maxLength={16} /><br /><br />

        <label>Имя:</label><br />
        <input type="text" name="firstName" value={form.firstName} onChange={handleChange} required /><br /><br />

        <label>Фамилия:</label><br />
        <input type="text" name="lastName" value={form.lastName} onChange={handleChange} required /><br /><br />

        <label>Отчество:</label><br />
        <input type="text" name="middleName" value={form.middleName} onChange={handleChange} /><br /><br />

        <label>Адрес:</label><br />
        <input type="text" name="address" value={form.address} onChange={handleChange} required /><br /><br />

        <label>Телефон:</label><br />
        <input type="tel" name="phoneNumber" pattern="[0-9]*" value={form.phoneNumber} onChange={handleChange} required /><br /><br />

        <label>Дата рождения:</label><br />
        <input type="date" name="dateOfBirth" value={form.dateOfBirth} onChange={handleChange} /><br /><br />

        <label>Дата приёма на работу:</label><br />
        <input type="date" name="dateOfHiring" value={form.dateOfHiring} onChange={handleChange} /><br /><br />

        <label>Должность:</label><br />
        <select name="position" value={form.position} onChange={handleChange}>
          {positions.map((pos) => (
            <option key={pos.ID} value={pos.ID}>
              {pos.Name}
            </option>
          ))}
        </select><br /><br />

        <input type="submit" value="Создать" />
      </form>
    </div>
  );
}
