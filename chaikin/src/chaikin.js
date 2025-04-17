const canvas = document.getElementById("canvas")

// Функция рисует линию по указанным точка.
function drawCurve(points, smooth) {
  // Создание новых точек для сглаживания кривой
  for (let i = 0; i < smooth; i++) {
    points = smoothing(points)
  }

  const context = canvas.getContext("2d");
  context.lineWidth = 10;
  context.strokeStyle = "#4663C9"

  context.beginPath();
  context.moveTo(points[0].x, points[0].y);
  for (let index = 1; index < points.length; index++) {
    context.lineTo(points[index].x, points[index].y);
  }
  context.stroke();
}

// Функция создает новые точки используя алгоритм Чайкина.
function smoothing(points) {
  const result = [points[0]]

  let index = 0
  const last_index = points.length - 1
  while (index < last_index) {
    const first = points[index]
    const second = points[index + 1]

    // Определение длины отрезка
    const dx = second.x - first.x
    const dy = second.y - first.y

    // 25% и 75% относительно длины полученного отрезка
    result.push({x: first.x + dx * 0.25, y: first.y + dy * 0.25})
    result.push({x: first.x + dx * 0.75, y: first.y + dy * 0.75})

    index++
  }

  result.push(points[last_index])
  return result
}