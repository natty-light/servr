export const ping = async () => {
  const res = await fetch('http://generate:3000/api/ping')  
  return res.json();
} 
