import React, { useState, useEffect } from "react";

const ConsultationForm = () => {
  const [clients, setClients] = useState([]);
  const [employees, setEmployees] = useState([]);
  const [form, setForm] = useState({
    dateofconsultation: "",
    timeofconsultation: "",
    client: "",
    employee: "",
    notes: "",
  });

  // Загружаем данные для select
  useEffect(() => {
    const fetchData = async () => {
      try {
        const [clientsRes, employeesRes] = await Promise.all([
          fetch("/api/clients"),
          fetch("/api/employees"),
        ]);

        if (clientsRes.ok) {
          setClients(await clientsRes.json());
        }
        if (employeesRes.ok) {
          setEmployees(await employeesRes.json());
        }
      } catch (err) {
        console.error("Ошибка загрузки данных:", err);
      }
    };
    fetchData();
  }, []);

  // Обновляем состояние формы
  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  // Отправляем форму
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch("/api/consultation/new", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify(form),
      });

      if (!res.ok) {
        const errorData = await res.json().catch(() => ({}));
        throw new Error(errorData.error || errorData.message || `HTTP error! status: ${res.status}`);
      }

      const data = await res.json();
      alert(data.message || "Консультация успешно создана");
      window.location.href = "/consultation";
    } catch (err) {
      console.error("Ошибка при создании консультации:", err);
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Добавление консультации</h2>
      <form onSubmit={handleSubmit}>
        <label>Дата консультации:</label><br />
        <input
          type="date"
          name="dateofconsultation"
          value={form.dateofconsultation}
          onChange={handleChange}
          required
        /><br /><br />

        <label>Время консультации:</label><br />
        <input
          type="time"
          name="timeofconsultation"
          value={form.timeofconsultation}
          onChange={handleChange}
          required
        /><br /><br />

        <label>Клиент:</label><br />
        <select
          name="client"
          value={form.client}
          onChange={handleChange}
          required
        >
          <option value="">-- Выберите клиента --</option>
          {clients.length > 0 ? (
            clients.map((c) => (
              <option key={c.id} value={c.id}>
                {c.firstname} {c.lastname}
              </option>
            ))
          ) : (
            <option value="" disabled>Нет доступных клиентов</option>
          )}
        </select><br /><br />

        <label>Примечания:</label><br />
        <input
          type="text"
          name="notes"
          value={form.notes}
          onChange={handleChange}
        /><br /><br />

        <label>Сотрудник:</label><br />
        <select
          name="employee"
          value={form.employee}
          onChange={handleChange}
          required
        >
          <option value="">-- Выберите сотрудника --</option>
          {employees.length > 0 ? (
            employees.map((e) => (
              <option key={e.id} value={e.id}>
                {e.firstName} {e.lastName}
              </option>
            ))
          ) : (
            <option value="" disabled>Нет доступных сотрудников</option>
          )}
        </select><br /><br />

        <button type="submit">Сохранить</button>
      </form>
    </div>
  );
};

export default ConsultationForm;
