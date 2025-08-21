import React from "react";
import { Link, useNavigate } from "react-router-dom";

const ToursList = ({ tours }) => {
  const navigate = useNavigate();

  const deleteTour = async (id) => {
    if (!window.confirm("Удалить тур?")) return;

    try {
      const res = await fetch(`/tours/${id}`, { method: "DELETE" });

      if (!res.ok) {
        const data = await res.json();
        throw new Error(data.error || "Ошибка при удалении");
      }

      alert("Тур удалён");
      // После удаления можно перезагрузить список
      navigate(0); // простая перезагрузка страницы
    } catch (err) {
      alert("Ошибка: " + err.message);
    }
  };

  return (
    <div>
      <h2>Список туров</h2>
      <table style={{ width: "100%", borderCollapse: "collapse" }} border="1" cellPadding="5" cellSpacing="0">
        <thead>
          <tr>
            <th>Название</th>
            <th>Рейтинг</th>
            <th>Отель</th>
            <th>Питание</th>
            <th>Город</th>
            <th>Страна</th>
            <th>Цена</th>
            <th>Действия</th>
          </tr>
        </thead>
        <tbody>
          {tours.length > 0 ? (
            tours.map((tour) => (
              <tr key={tour.ID}>
                <td>{tour.Name}</td>
                <td>{tour.Rating}</td>
                <td>{tour.Hotel}</td>
                <td>{tour.Nutrition}</td>
                <td>{tour.City}</td>
                <td>{tour.Country}</td>
                <td>{tour.Price} ₽</td>
                <td>
                  <Link to={`/tours/edit/${tour.ID}`}>✏️ Редактировать</Link>{" "}
                  {/* <button onClick={() => deleteTour(tour.ID)}>🗑️ Удалить</button> */}
                </td>
              </tr>
            ))
          ) : (
            <tr>
              <td colSpan="8" style={{ textAlign: "center" }}>
                Нет доступных туров
              </td>
            </tr>
          )}
        </tbody>
      </table>

      <br />
      <Link to="/tours/new">➕ Добавить тур</Link>
    </div>
  );
};

export default ToursList;
