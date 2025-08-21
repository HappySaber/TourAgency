import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

export default function ConsultationList() {
  const [consultations, setConsultations] = useState([]);
  const navigate = useNavigate();

  useEffect(() => {
    loadConsultations();
  }, []);

  async function loadConsultations() {
    try {
      const res = await fetch("/consultation");
      if (!res.ok) throw new Error("Ошибка загрузки консультаций");
      const data = await res.json();
      setConsultations(data);
    } catch (err) {
      alert("Ошибка: " + err.message);
    }
  }

  async function deleteConsultation(id) {
    if (!window.confirm("Удалить консультацию?")) return;

    try {
      const res = await fetch(`/consultation/${id}`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
      });

      if (res.ok) {
        alert("Консультация удалена");
        setConsultations((prev) => prev.filter((c) => c.id !== id));
      } else {
        const data = await res.json();
        alert(data.error || "Ошибка при удалении");
      }
    } catch (err) {
      alert("Ошибка: " + err.message);
    }
  }

  return (
    <div className="p-4">
      <h2 className="text-xl mb-4">Список консультаций</h2>

      <table border="1" cellPadding="5" cellSpacing="0" className="w-full">
        <thead>
          <tr>
            <th>Дата консультации</th>
            <th>Время консультации</th>
            <th>Клиент</th>
            <th>Сотрудник</th>
            <th>Примечания</th>
            <th>Действия</th>
          </tr>
        </thead>
        <tbody>
          {consultations.map((c) => (
            <tr key={c.id}>
              <td>{new Date(c.dateofconsultation).toLocaleDateString()}</td>
              <td>{c.timeofconsultation}</td>
              <td>
                {c.client?.firstname} {c.client?.lastname}
              </td>
              <td>
                {c.employee?.firstname} {c.employee?.lastname}
              </td>
              <td>{c.notes}</td>
              <td>
                <button onClick={() => navigate(`/consultation/edit/${c.id}`)}>
                  ✏️ Редактировать
                </button>
                <button onClick={() => navigate(`/consultation/${c.id}/tours`)}>
                  🧳 Туры
                </button>
                <button
                  onClick={() => navigate(`/consultation/${c.id}/services`)}
                >
                  🛠 Услуги
                </button>
                <button onClick={() => deleteConsultation(c.id)}>🗑️ Удалить</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>

      <br />
      <button onClick={() => navigate("/consultation/new")}>
        ➕ Добавить консультацию
      </button>
    </div>
  );
}
