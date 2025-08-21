import React, { useState } from "react";

const EditConsultationForm = ({ consultation, clients, employees }) => {
  const [formData, setFormData] = useState({
    id: consultation.id,
    dateofconsultation: consultation.dateofconsultation,
    timeofconsultation: consultation.timeofconsultation,
    client: consultation.clientId || "",
    employee: consultation.employeeId || "",
    notes: consultation.notes || "",
  });

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: value,
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const res = await fetch(`/consultation/edit/${formData.id}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      if (res.ok) {
        alert("Консультация обновлена");
        window.location.href = "/consultation";
      } else {
        const data = await res.json();
        alert(data.error || "Ошибка при обновлении");
      }
    } catch (err) {
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Редактирование консультации</h2>
      <form onSubmit={handleSubmit}>
        <label>Дата консультации:</label>
        <br />
        <input
          type="date"
          name="dateofconsultation"
          value={formData.dateofconsultation}
          onChange={handleChange}
          required
        />
        <br />
        <br />

        <label>Время консультации:</label>
        <br />
        <input
          type="time"
          name="timeofconsultation"
          value={formData.timeofconsultation}
          onChange={handleChange}
          required
        />
        <br />
        <br />

        <label>Клиент:</label>
        <br />
        <select
          name="client"
          value={formData.client}
          onChange={handleChange}
          required
        >
          <option value="">-- Выберите клиента --</option>
          {clients.map((c) => (
            <option key={c.id} value={c.id}>
              {c.firstname} {c.lastname}
            </option>
          ))}
        </select>
        <br />
        <br />

        <label>Примечания:</label>
        <br />
        <input
          type="text"
          name="notes"
          value={formData.notes}
          onChange={handleChange}
        />
        <br />
        <br />

        <label>Сотрудник:</label>
        <br />
        <select
          name="employee"
          value={formData.employee}
          onChange={handleChange}
          required
        >
          <option value="">-- Выберите сотрудника --</option>
          {employees.map((e) => (
            <option key={e.id} value={e.id}>
              {e.firstName} {e.lastName}
            </option>
          ))}
        </select>
        <br />
        <br />

        <button type="submit">Сохранить изменения</button>
      </form>
    </div>
  );
};

export default EditConsultationForm;
