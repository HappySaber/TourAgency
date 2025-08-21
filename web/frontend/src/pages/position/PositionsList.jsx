// src/pages/position/PositionList.jsx
import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const API_URL = "http://localhost:8080/api";

export default function PositionList() {
  const [positions, setPositions] = useState([]);
  const [loading, setLoading] = useState(true);

  // Загрузка списка должностей
  useEffect(() => {
    fetch(`${API_URL}/position`, {
      credentials: "include", // если у тебя авторизация по cookie
    })
      .then((res) => res.json())
      .then((data) => {
        setPositions(data);
        setLoading(false);
      })
      .catch((err) => {
        console.error("Ошибка загрузки:", err);
        setLoading(false);
      });
  }, []);

  // Удаление должности
  const deletePosition = async (id) => {
    if (!window.confirm("Удалить должность?")) return;

    try {
      const res = await fetch(`${API_URL}/position/${id}`, {
        method: "DELETE",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
      });

      if (res.ok) {
        setPositions(positions.filter((p) => p.ID !== id));
      } else {
        const data = await res.json();
        alert(data.error || "Ошибка при удалении");
      }
    } catch (error) {
      console.error("Ошибка удаления:", error);
      alert("Ошибка сети");
    }
  };

  if (loading) {
    return <p>Загрузка...</p>;
  }

  return (
    <div>
      <h2>Список должностей</h2>
      <Link to="/position/create" className="btn btn-primary">
        Добавить должность
      </Link>
      <table border="1" cellPadding="8" style={{ marginTop: "10px" }}>
        <thead>
          <tr>
            <th>ID</th>
            <th>Название должности</th>
            <th>Действия</th>
          </tr>
        </thead>
        <tbody>
          {positions.length > 0 ? (
            positions.map((pos) => (
              <tr key={pos.ID}>
                <td>{pos.ID}</td>
                <td>{pos.name}</td>
                <td>
                  <Link to={`/position/edit/${pos.ID}`}>Редактировать</Link>{" "}
                  |{" "}
                  <button
                    onClick={() => deletePosition(pos.ID)}
                    style={{ color: "red", border: "none", background: "none", cursor: "pointer" }}
                  >
                    Удалить
                  </button>
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan="3">Нет должностей</td>
            </tr>
          )}
        </tbody>
      </table>
    </div>
  );
}
