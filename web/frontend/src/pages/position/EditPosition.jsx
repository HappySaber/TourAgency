import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const EditPosition = ({ position, csrfToken }) => {
  const [name, setName] = useState(position.Name);
  const [salary, setSalary] = useState(position.Salary);
  const [responsibilities, setResponsibilities] = useState(position.Responsibilities);
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    const updatedPosition = {
      id: position.ID,
      name,
      salary,
      responsibilities,
    };

    try {
      const res = await fetch(`/position/edit/${position.ID}`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
          "X-CSRF-Token": csrfToken,
        },
        body: JSON.stringify(updatedPosition),
      });

      const data = await res.json();

      if (!res.ok) {
        throw new Error(data.error || data.message || "Ошибка при обновлении");
      }

      alert("Должность обновлена");
      navigate("/position");
    } catch (err) {
      console.error("Ошибка при обновлении должности:", err);
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Редактировать должность</h2>
      <form onSubmit={handleSubmit}>
        <label>Название:</label>
        <br />
        <input
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          required
        />
        <br />
        <br />

        <label>Зарплата:</label>
        <br />
        <input
          type="text"
          value={salary}
          onChange={(e) => setSalary(e.target.value)}
          required
        />
        <br />
        <br />

        <label>Обязанности:</label>
        <br />
        <input
          type="text"
          value={responsibilities}
          onChange={(e) => setResponsibilities(e.target.value)}
          required
        />
        <br />
        <br />

        <button type="submit">Сохранить изменения</button>
      </form>
    </div>
  );
};

export default EditPosition;
