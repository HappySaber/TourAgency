import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const CreatePosition = () => {
  const [name, setName] = useState("");
  const [salary, setSalary] = useState("");
  const [responsibilities, setResponsibilities] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();

    const position = { name, salary, responsibilities };

    try {
      const res = await fetch("/position/new", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Accept: "application/json",
        },
        body: JSON.stringify(position),
      });

      const data = await res.json();

      if (!res.ok) {
        throw new Error(data.error || data.message || "Неизвестная ошибка");
      }

      alert(data.message || "Должность успешно создана");
      navigate("/position");
    } catch (err) {
      console.error("Ошибка при создании должности:", err);
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Добавить должность</h2>
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

        <button type="submit">Сохранить</button>
      </form>
    </div>
  );
};

export default CreatePosition;
