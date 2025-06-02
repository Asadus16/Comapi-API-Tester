import React from 'react';

const TestResults = ({ results }) => {
  if (!results) return null;

  const { suite_name, total_tests, passed_tests, failed_tests, results: testResults } = results;
  const successRate = total_tests > 0 ? Math.round((passed_tests / total_tests) * 100) : 0;

  return (
    <div className="space-y-6">
      {/* Results Header */}
      <div className="bg-dark-surface rounded-xl border border-dark-border p-6">
        <div className="flex items-center space-x-3 mb-6">
          <div className="w-8 h-8 bg-green-500/20 rounded-lg flex items-center justify-center">
            <span className="text-green-400">📊</span>
          </div>
          <h2 className="text-xl font-bold text-green-400">
            Test Results
          </h2>
        </div>
        
        {/* Stats Grid */}
        <div className="grid grid-cols-2 lg:grid-cols-54 gap-10">

          <StatCard
            label="Passed"
            value={passed_tests}
            color="green"
            icon="✅"
          />
          <StatCard
            label="Failed"
            value={failed_tests}
            color="red"
            icon="❌"
          />
 
        </div>

        {/* Success Rate Bar */}
        <div className="mt-6">
          <div className="flex justify-between text-sm text-dark-text-secondary mb-2">
            <span>Overall Success Rate</span>
            <span>{successRate}%</span>
          </div>
          <div className="w-full bg-dark-bg rounded-full h-3">
            <div
              className={`h-3 rounded-full transition-all duration-1000 ${
                successRate >= 80 ? 'bg-green-500' : 
                successRate >= 50 ? 'bg-yellow-500' : 'bg-red-500'
              }`}
              style={{ width: `${successRate}%` }}
            ></div>
          </div>
        </div>
      </div>

      {/* Individual Test Results */}
      <div className="space-y-4">
        {testResults.map((result, index) => (
          <TestResultCard key={index} result={result} index={index} />
        ))}
      </div>
    </div>
  );
};

// Stat Card Component
const StatCard = ({ label, value, color, icon }) => {
  const colorClasses = {
    blue: 'bg-blue-500/20 border-blue-500/30 text-blue-400',
    green: 'bg-green-500/20 border-green-500/30 text-green-400',
    red: 'bg-red-500/20 border-red-500/30 text-red-400',
    yellow: 'bg-yellow-500/20 border-yellow-500/30 text-yellow-400',
  };

  return (
    <div className={`border rounded-lg p-4 ${colorClasses[color]}`}>
      <div className="flex items-center justify-between mb-2">
        <span className="text-2xl">{icon}</span>
        <span className="text-2xl font-bold">{value}</span>
      </div>
      <p className="text-sm opacity-80">{label}</p>
    </div>
  );
};

// Individual Test Result Card
const TestResultCard = ({ result, index }) => {
  const { test_name, status, duration, request, response, assertions = [], error } = result;
  const isPassed = status === 'PASS';
  const durationMs = Math.round(duration / 1000000);

  const methodColors = {
    GET: 'bg-green-500/20 text-green-400 border-green-500/30',
    POST: 'bg-blue-500/20 text-blue-400 border-blue-500/30',
    PUT: 'bg-yellow-500/20 text-yellow-400 border-yellow-500/30',
    DELETE: 'bg-red-500/20 text-red-400 border-red-500/30',
    PATCH: 'bg-purple-500/20 text-purple-400 border-purple-500/30'
  };

  return (
    <div className={`
      bg-dark-surface rounded-xl border p-6 transition-all
      ${isPassed 
        ? 'border-green-500/30 bg-green-500/5' 
        : 'border-red-500/30 bg-red-500/5'
      }
    `}>
      {/* Test Header */}
      <div className="flex items-start justify-between mb-4">
        <div className="flex-1">
          <div className="flex items-center space-x-3 mb-2">
            <span className="text-dark-text-secondary font-mono text-sm">
              #{index + 1}
            </span>
            <h3 className="text-lg font-semibold text-dark-text">
              {test_name}
            </h3>
          </div>
          
          <div className="flex items-center space-x-3 text-sm">
            <span className={`px-2 py-1 rounded border font-mono font-bold ${methodColors[request.method]}`}>
              {request.method}
            </span>
            <span className="text-dark-text-secondary font-mono">
              {request.url}
            </span>
          </div>
        </div>
        
        <div className="flex flex-col items-end space-y-2">
          <div className={`
            px-3 py-1 rounded-lg font-bold text-sm flex items-center space-x-2
            ${isPassed 
              ? 'bg-green-500/20 text-green-400 border border-green-500/30' 
              : 'bg-red-500/20 text-red-400 border border-red-500/30'
            }
          `}>
            <span>{isPassed ? '✅' : '❌'}</span>
            <span>{status}</span>
          </div>
          
          <div className="text-sm text-dark-text-secondary">
            {durationMs}ms
          </div>
        </div>
      </div>

      {/* Error Message */}
      {error && (
        <div className="bg-red-900/20 border border-red-500/30 rounded-lg p-4 mb-4">
          <div className="flex items-start space-x-2">
            <span className="text-red-400 text-lg">⚠️</span>
            <div>
              <h4 className="text-red-400 font-semibold mb-1">Error</h4>
              <p className="text-red-300 text-sm">{error}</p>
            </div>
          </div>
        </div>
      )}

      {/* Response Info */}
      <div className="grid grid-cols-2 gap-4 mb-4">
        <div className="bg-dark-bg rounded-lg p-3">
          <div className="text-xs text-dark-text-secondary uppercase tracking-wide mb-1">
            Status Code
          </div>
          <div className={`text-lg font-bold ${
            response.status_code >= 200 && response.status_code < 300 ? 'text-green-400' :
            response.status_code >= 400 ? 'text-red-400' : 'text-yellow-400'
          }`}>
            {response.status_code}
          </div>
        </div>
        
        <div className="bg-dark-bg rounded-lg p-3">
          <div className="text-xs text-dark-text-secondary uppercase tracking-wide mb-1">
            Response Size
          </div>
          <div className="text-lg font-bold text-dark-text">
            {formatBytes(response.size)}
          </div>
        </div>
      </div>

      {/* Assertions */}
      {assertions.length > 0 && (
        <div className="space-y-3">
          <h4 className="text-sm font-semibold text-dark-text uppercase tracking-wide">
            Assertions
          </h4>
          <div className="space-y-2">
            {assertions.map((assertion, assertionIndex) => (
              <div
                key={assertionIndex}
                className={`
                  flex items-center space-x-3 p-3 rounded-lg border text-sm
                  ${assertion.passed
                    ? 'bg-green-500/10 border-green-500/20 text-green-400'
                    : 'bg-red-500/10 border-red-500/20 text-red-400'
                  }
                `}
              >
                <span className="text-lg">
                  {assertion.passed ? '✅' : '❌'}
                </span>
                <span className="font-semibold min-w-20">
                  {assertion.type.toUpperCase()}
                </span>
                <span className="flex-1">
                  {assertion.message}
                </span>
              </div>
            ))}
          </div>
        </div>
      )}

      {/* Response Body (for all tests, not just failed ones) */}
      {response.body && (
        <details className="mt-4">
          <summary className="text-neon-orange cursor-pointer hover:text-neon-orange/80 font-medium flex items-center space-x-2">
            <span>📄</span>
            <span>View Full JSON Response</span>
            <span className="text-xs text-dark-text-secondary">({formatBytes(response.size)})</span>
          </summary>
          <div className="mt-3 bg-dark-bg rounded-lg p-4 border border-dark-border">
            <div className="mb-2 flex justify-between items-center">
              <span className="text-xs text-dark-text-secondary uppercase tracking-wide">Response Body</span>
              <button 
                onClick={() => copyToClipboard(response.body)}
                className="text-xs bg-neon-orange/20 text-neon-orange px-2 py-1 rounded hover:bg-neon-orange/30 transition-colors"
              >
                📋 Copy
              </button>
            </div>
            <pre className="text-xs text-dark-text overflow-x-auto font-mono max-h-96 overflow-y-auto">
              {formatJSON(response.body)}
            </pre>
          </div>
        </details>
      )}
    </div>
  );
};

// Helper function to format bytes
const formatBytes = (bytes) => {
  if (bytes === 0) return '0 B';
  const k = 1024;
  const sizes = ['B', 'KB', 'MB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));
  return parseFloat((bytes / Math.pow(k, i)).toFixed(1)) + ' ' + sizes[i];
};

// Helper function to format JSON
const formatJSON = (jsonString) => {
  try {
    const parsed = JSON.parse(jsonString);
    return JSON.stringify(parsed, null, 2);
  } catch (e) {
    return jsonString; // Return as-is if not valid JSON
  }
};

// Helper function to copy to clipboard
const copyToClipboard = (text) => {
  navigator.clipboard.writeText(text).then(() => {
    // Could add a toast notification here
    console.log('Copied to clipboard!');
  }).catch(err => {
    console.error('Failed to copy: ', err);
  });
};

export default TestResults;