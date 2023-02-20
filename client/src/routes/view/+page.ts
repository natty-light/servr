import type { generateRequest, schoolOption } from '../../types';

export const _postGenerateRequest = async (request: generateRequest) => {
  const res = await fetch(`http://localhost:3000/api/generate`, {
    method: 'POST',
    body: prepareRequestBody(request.schools),
    headers: {
      'Authorization': `Bearer ${request.tokenString}`
    }
  })

  // Turn stream into blob
  return await res.blob()

};

const prepareRequestBody = (schools: schoolOption[]) => {
  const schoolArray: string[] = schools.map( (school) => school.text)
  const body = {
    schoolArray
  }
  return JSON.stringify(body);
}