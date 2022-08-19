import { useEffect, useState } from 'react';
import { Chessboard } from 'react-chessboard';

const uuid = "5198a647-aede-4e85-8294-37b6f9526a9e";
const ws = new WebSocket("ws://" + window.location.hostname + ":8085/ws");

export default function App() {
  const [position, setPosition] = useState<any>({});

  useEffect(() => {
    fetch("http://" + window.location.hostname + ":8085/api/v1/session/" + uuid)
      .then((response) => response.json())
      .then((data) => setPosition(data.position));
  }, []);

  ws.onmessage = function (event) {
    const json = JSON.parse(event.data);
    if (position !== json.position) {
      setPosition(json.position)
    }
  };

  function onDrop(sourceSquare: any, targetSquare: any, piece: string) {
    console.log(sourceSquare, targetSquare, piece)
    if (position && sourceSquare in position && position[sourceSquare] === piece) {
      let newPosition: any = position
      delete newPosition[sourceSquare]
      newPosition[targetSquare] = piece
      setPosition(newPosition)

      ws.send(JSON.stringify({
        uuid: uuid,
        position: newPosition
      }))
    }
    return false
  }

  return (
    <div>
      <Chessboard id={0} position={position} onPieceDrop={onDrop} />
    </div>
  );
}
