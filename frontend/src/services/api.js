// API service to communicate with Go backend

const API_BASE_URL = 'http://localhost:8080/api/v1';

// Run tests endpoint
export const runTests = async (testSuite) => {
  try {
    const response = await fetch(`${API_BASE_URL}/tests/run`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ test_suite: testSuite }),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.error || `HTTP ${response.status}: ${response.statusText}`);
    }

    return await response.json();
  } catch (error) {
    if (error.name === 'TypeError' && error.message.includes('fetch')) {
      throw new Error('Unable to connect to Comapi server. Make sure the Go backend is running on port 8080.');
    }
    throw error;
  }
};

// Validate tests endpoint
export const validateTests = async (testSuite) => {
  try {
    const response = await fetch(`${API_BASE_URL}/tests/validate`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ test_suite: testSuite }),
    });

    return await response.json();
  } catch (error) {
    throw new Error('Failed to validate tests: ' + error.message);
  }
};

// Health check endpoint
export const healthCheck = async () => {
  try {
    const response = await fetch(`${API_BASE_URL}/health`);
    return await response.json();
  } catch (error) {
    throw new Error('Backend server is not responding');
  }
};