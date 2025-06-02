import React from 'react';

const TestConfiguration = ({ testSuite, onChange, onRunTests, loading }) => {
  // Since we only have one test, we'll work with the first test in the array
  const currentTest = testSuite.tests[0] || {
    name: '',
    method: 'GET',
    url: '', // Single complete URL field
    headers: {},
    body: '',
    assertions: [{ type: 'status', expected: 200 }]
  };

  const updateTest = (field, value) => {
    const updatedTest = { ...currentTest, [field]: value };
    onChange({
      ...testSuite,
      name: "Single Test", // Auto-set suite name
      base_url: "", // Not needed anymore
      tests: [updatedTest]
    });
  };

  const canRunTest = currentTest.name && currentTest.url;

  const methodColors = {
    GET: 'bg-green-500/20 text-green-400 border-green-500/30',
    POST: 'bg-blue-500/20 text-blue-400 border-blue-500/30',
    PUT: 'bg-yellow-500/20 text-yellow-400 border-yellow-500/30',
    DELETE: 'bg-red-500/20 text-red-400 border-red-500/30',
    PATCH: 'bg-purple-500/20 text-purple-400 border-purple-500/30'
  };

  return (
    <div className="space-y-6">
      {/* Single Test Configuration */}
      <div className="bg-dark-surface rounded-xl border border-dark-border p-6">
        <div className="flex items-center space-x-3 mb-6">
          <div className="w-8 h-8 bg-neon-orange/20 rounded-lg flex items-center justify-center">
            <span className="text-neon-orange">🧪</span>
          </div>
          <h2 className="text-xl font-bold text-neon-orange">
            API Test
          </h2>
        </div>
        
        <div className="space-y-6">
          {/* Test Name */}
          <div>
            <label className="block text-sm font-medium text-dark-text mb-2">
              Test Name
            </label>
            <input
              type="text"
              placeholder="My API Test"
              value={currentTest.name}
              onChange={(e) => updateTest('name', e.target.value)}
              className="w-full bg-dark-bg border border-dark-border rounded-lg px-4 py-3 text-dark-text placeholder-dark-text-secondary focus:border-neon-orange focus:ring-1 focus:ring-neon-orange transition-colors"
            />
          </div>

          {/* Method and URL */}
          <div>
            <label className="block text-sm font-medium text-dark-text mb-2">
              Request
            </label>
            <div className="flex space-x-3">
              <select
                value={currentTest.method}
                onChange={(e) => updateTest('method', e.target.value)}
                className={`px-4 py-3 rounded-lg border font-mono font-bold text-sm ${methodColors[currentTest.method]} bg-dark-surface focus:outline-none focus:ring-2 focus:ring-neon-orange/50`}
              >
                <option value="GET">GET</option>
                <option value="POST">POST</option>
                <option value="PUT">PUT</option>
                <option value="DELETE">DELETE</option>
                <option value="PATCH">PATCH</option>
              </select>
              
              <input
                type="text"
                placeholder="https://api.example.com/users/1"
                value={currentTest.url || ''}
                onChange={(e) => updateTest('url', e.target.value)}
                className="flex-1 bg-dark-bg border border-dark-border rounded-lg px-4 py-3 text-dark-text placeholder-dark-text-secondary focus:border-neon-orange focus:ring-1 focus:ring-neon-orange transition-colors font-mono"
              />
            </div>
          </div>

          {/* Request Headers */}
          <div>
            <label className="block text-sm font-medium text-dark-text mb-2">
              Headers (Optional)
            </label>
            <div className="grid grid-cols-1 gap-3">
              <div className="flex space-x-3">
                <input
                  type="text"
                  placeholder="Content-Type"
                  className="flex-1 bg-dark-bg border border-dark-border rounded-lg px-4 py-2 text-dark-text placeholder-dark-text-secondary focus:border-neon-orange focus:ring-1 focus:ring-neon-orange transition-colors text-sm"
                />
                <input
                  type="text"
                  placeholder="application/json"
                  className="flex-1 bg-dark-bg border border-dark-border rounded-lg px-4 py-2 text-dark-text placeholder-dark-text-secondary focus:border-neon-orange focus:ring-1 focus:ring-neon-orange transition-colors text-sm"
                />
              </div>
              <div className="flex space-x-3">
                <input
                  type="text"
                  placeholder="Authorization"
                  className="flex-1 bg-dark-bg border border-dark-border rounded-lg px-4 py-2 text-dark-text placeholder-dark-text-secondary focus:border-neon-orange focus:ring-1 focus:ring-neon-orange transition-colors text-sm"
                />
                <input
                  type="text"
                  placeholder="Bearer token..."
                  className="flex-1 bg-dark-bg border border-dark-border rounded-lg px-4 py-2 text-dark-text placeholder-dark-text-secondary focus:border-neon-orange focus:ring-1 focus:ring-neon-orange transition-colors text-sm"
                />
              </div>
            </div>
          </div>

          {/* Request Body for non-GET requests */}
          {currentTest.method !== 'GET' && (
            <div>
              <label className="block text-sm font-medium text-dark-text mb-2">
                Request Body (JSON)
              </label>
              <textarea
                placeholder='{"key": "value"}'
                value={currentTest.body || ''}
                onChange={(e) => updateTest('body', e.target.value)}
                className="w-full bg-dark-bg border border-dark-border rounded-lg px-4 py-3 text-dark-text placeholder-dark-text-secondary focus:border-neon-orange focus:ring-1 focus:ring-neon-orange transition-colors font-mono text-sm"
                rows={6}
              />
            </div>
          )}

          {/* Assertions Preview */}
          <div className="bg-neon-orange/10 border border-neon-orange/20 rounded-lg p-4">
            <div className="flex items-center space-x-3 mb-3">
              <span className="text-neon-orange text-lg">✓</span>
              <span className="text-lg font-semibold text-neon-orange">
                Assertions
              </span>
            </div>
            <div className="space-y-2 text-sm">
              <div className="flex items-center space-x-2">
                <span className="w-2 h-2 bg-green-400 rounded-full"></span>
                <span className="text-green-400">Status Code = 200</span>
              </div>
              <div className="flex items-center space-x-2">
                <span className="w-2 h-2 bg-blue-400 rounded-full"></span>
                <span className="text-blue-400">Response Time &lt; 5000ms</span>
              </div>
              <div className="flex items-center space-x-2">
                <span className="w-2 h-2 bg-purple-400 rounded-full"></span>
                <span className="text-purple-400">Content-Type validation</span>
              </div>
            </div>
          </div>
        </div>
      </div>

      {/* Run Test Section */}
      <div className="bg-dark-surface rounded-xl border border-dark-border p-8">
        <div className="text-center">
          <button
            onClick={onRunTests}
            disabled={!canRunTest || loading}
            className={`
              px-12 py-5 rounded-xl font-bold text-xl transition-all duration-300
              ${canRunTest && !loading
                ? 'bg-neon-orange hover:bg-neon-orange/90 text-white shadow-neon hover:shadow-neon transform hover:scale-105 hover:-translate-y-1'
                : 'bg-gray-600 text-gray-400 cursor-not-allowed'
              }
            `}
          >
            {loading ? (
              <div className="flex items-center space-x-4">
                <div className="w-6 h-6 border-3 border-white/30 border-t-white rounded-full animate-spin"></div>
                <span>Testing API...</span>
              </div>
            ) : (
              <div className="flex items-center space-x-4">
                <span className="text-2xl">🚀</span>
                <span>RUN TEST</span>
              </div>
            )}
          </button>
          
          {!canRunTest && !loading && (
            <div className="mt-4 space-y-2">
              <p className="text-dark-text-secondary text-sm">
                Please provide a test name and complete URL
              </p>
              <div className="text-xs text-dark-text-secondary space-y-1">
                {!currentTest.name && <div>• Add a test name</div>}
                {!currentTest.url && <div>• Add a complete URL (e.g., https://api.example.com/users)</div>}
              </div>
            </div>
          )}
        </div>
      </div>

      {/* Quick Examples */}
      <div className="bg-dark-surface/50 rounded-xl border border-dark-border/50 p-4">
        <h3 className="text-sm font-semibold text-dark-text mb-3">Quick Examples</h3>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-3 text-xs">
          <button 
            onClick={() => {
              updateTest('name', 'Get GitHub User');
              updateTest('method', 'GET');
              updateTest('url', 'https://api.github.com/users/octocat');
            }}
            className="text-left p-3 bg-dark-bg rounded-lg border border-dark-border hover:border-neon-orange/50 transition-colors"
          >
            <div className="font-mono text-green-400">GET</div>
            <div className="text-dark-text">GitHub User API</div>
            <div className="text-dark-text-secondary">api.github.com/users/octocat</div>
          </button>
          
          <button 
            onClick={() => {
              updateTest('name', 'JSONPlaceholder Post');
              updateTest('method', 'GET');
              updateTest('url', 'https://jsonplaceholder.typicode.com/posts/1');
            }}
            className="text-left p-3 bg-dark-bg rounded-lg border border-dark-border hover:border-neon-orange/50 transition-colors"
          >
            <div className="font-mono text-green-400">GET</div>
            <div className="text-dark-text">Test JSON API</div>
            <div className="text-dark-text-secondary">jsonplaceholder.typicode.com/posts/1</div>
          </button>
        </div>
      </div>
    </div>
  );
};

export default TestConfiguration;