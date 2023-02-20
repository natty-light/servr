import type { RequestHandler } from '@sveltejs/kit';
import type { generateRequest, schoolOption } from '../../../types';

export const POST = (async ({request}) => {
  const req: generateRequest = await request.json();
  const res = await fetch(`http://localhost:3000/api/generate`, {
    method: 'POST',
    body: prepareRequestBody(req.schools),
    headers: {
      'Authorization': `Bearer ${req.tokenString}`
    }
  })

  // Turn stream into blob
  const response = await res.arrayBuffer()
  const binaryString = Buffer.from(response).toString('binary');
  const bytes = new Uint8Array(binaryString.length);
  for (let i = 0; i < binaryString.length; i++) {
    bytes[i] = binaryString.charCodeAt(i);
  }

  const blob = new Blob([bytes], {type: 'application/pdf'});

  return new Response(blob);

}) satisfies RequestHandler;

const prepareRequestBody = (schools: schoolOption[]) => {
  const schoolArray: string[] = schools.map( (school) => school.text)
  const body = {
    schoolArray
  }
  return JSON.stringify(body);
}