import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const ProvidersList = () => {
  const [providers, setProviders] = useState([]);
  const navigate = useNavigate();

  // Загрузка списка провайдеров с сервера
  useEffect(() => {
    fetch("/provider") // предполагаемая API точка для получения списка
      .then((res) => res.json())
      .then((data) => setProviders(data))
      .catch((err) => console.error("Ошибка при загрузке провайдеров:", err));
  }, []);

  const deleteProvider = async (id) => {
    if (!window.confirm("Удалить провайдера?")) return;

    try {
      const res = await fetch(`/provider/delete/${id}`, { method: "DELETE" });
      if (!res.ok) {
        const data = await res.json();
        throw new Error(data.error || "Ошибка при удалении");
      }

      // Обновляем локальный список после удаления
      setProviders(providers.filter((p) => p.ID !== id));
      alert("Провайдер удален");
    } catch (err) {
      console.error(err);
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Список провайдеров</h2>
      <table border="1" cellPadding="5" cellSpacing="0">
        <thead>
          <tr>
            <th>Название</th>
            <th>Обращаться К</th>
            <th>Адрес</th>
            <th>Email</th>
            <th>Телефон</th>
            <th>Действия</th>
          </tr>
        </thead>
        <tbody>
          {providers.map((p) => (
            <tr key={p.ID}>
              <td>{p.Name}</td>
              <td>{p.Addressto}</td>
              <td>{p.Address}</td>
              <td>{p.Email}</td>
              <td>{p.PhoneNumber}</td>
              <td>
                <button onClick={() => navigate(`/provider/edit/${p.ID}`)}>✏️ Редактировать</button>
                <button onClick={() => deleteProvider(p.ID)}>🗑️ Удалить</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      <br />
      <button onClick={() => navigate("/provider/new")}>➕ Добавить провайдера</button>
    </div>
  );
};

export default ProvidersList;
