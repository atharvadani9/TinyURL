import { addDays, format } from "date-fns";

function App() {
  return <>Tomorrow's Date: {format(addDays(new Date(), 1), "do MMMM yyyy")}</>;
}

export default App;
