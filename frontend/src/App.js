import React, { useState } from 'react';
import TestConfiguration from './components/TestConfiguration';
import TestResults from './components/TestResults';
import LoadingSpinner from './components/LoadingSpinner';
import { runTests } from './services/api';
import Header from './components/header';

const App = () => {
  const [testSuite, setTestSuite] = useState({
    name: '',
    base_url: '',
    headers: {},
    tests: []
  });
  
  const [results, setResults] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleRunTests = async () => {
    setLoading(true);
    setError(null);
    setResults(null);
    
    try {
      const data = await runTests(testSuite);
      setResults(data);
    } catch (err) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-dark-bg">
      {/* Header */}
    <Header/>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 py-6">
        <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
          {/* Left Panel - Configuration */}
          <div className="space-y-6">
            <TestConfiguration
              testSuite={testSuite}
              onChange={setTestSuite}
              onRunTests={handleRunTests}
              loading={loading}
            />
          </div>

          {/* Right Panel - Results */}
          <div className="space-y-6">
            {loading && (
              <div className="bg-dark-surface rounded-xl border border-dark-border p-8">
                <LoadingSpinner message="Running your API tests..." />
              </div>
            )}
            
            {error && (
              <div className="bg-red-900/20 border border-red-500/30 rounded-xl p-6">
                <div className="flex items-start space-x-3">
                  <div className="w-6 h-6 bg-red-500 rounded-full flex items-center justify-center flex-shrink-0 mt-0.5">
                    <span className="text-white text-sm">!</span>
                  </div>
                  <div>
                    <h3 className="text-red-400 font-semibold text-lg mb-2">
                      Connection Failed
                    </h3>
                    <p className="text-red-300 mb-4">{error}</p>
                    <details className="text-sm">
                      <summary className="text-red-400 cursor-pointer hover:text-red-300">
                        Troubleshooting Steps
                      </summary>
                      <ul className="mt-2 space-y-1 text-red-200 ml-4">
                        <li>• Make sure Go backend is running: <code className="bg-red-800/30 px-1 rounded">go run main.go server</code></li>
                        <li>• Check server is accessible at <span className="text-neon-orange">http://localhost:8080</span></li>
                        <li>• Verify your test configuration is valid</li>
                      </ul>
                    </details>
                  </div>
                </div>
              </div>
            )}
            
            {results && <TestResults results={results} />}
            
            {!loading && !error && !results && (
              <div className="bg-dark-surface rounded-xl border border-dark-border p-12 text-center">
                <div className="w-16 h-16 bg-dark-border rounded-full flex items-center justify-center mx-auto mb-4">
                  <span className="text-2xl">📊</span>
                </div>
                <h3 className="text-xl font-semibold text-dark-text mb-2">
                  Ready to Test
                </h3>
                <p className="text-dark-text-secondary">
                  Configure your tests and click run to see results here
                </p>
              </div>
            )}
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer className="mt-12 border-t border-dark-border">
        <div className="max-w-7xl mx-auto px-4 py-6 text-center">
          <p className="text-dark-text-secondary">
            Built with <span className="text-neon-orange">Go + React</span>
          </p>
        </div>
      </footer>
    </div>
  );
};

export default App;
