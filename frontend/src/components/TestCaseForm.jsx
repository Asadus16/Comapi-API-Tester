import React from 'react';
import AssertionList from '../Assertions/AssertionList';

const TestCaseForm = ({ test, onChange, onDelete }) => {
  const updateTest = (field, value) => {
    onChange({ ...test, [field]: value });
  };

  return (
    <div className="test-case-form">
      <div className="test-header">
        <input
          type="text"
          value={test.name}
          onChange={(e) => updateTest('name', e.target.value)}
          placeholder="Test name"
          className="test-name-input"
        />
        
        <button onClick={onDelete} className="delete-btn">
          🗑️ Delete
        </button>
      </div>
      
      <div className="test-config">
        <select
          value={test.method}
          onChange={(e) => updateTest('method', e.target.value)}
        >
          <option value="GET">GET</option>
          <option value="POST">POST</option>
          <option value="PUT">PUT</option>
          <option value="DELETE">DELETE</option>
        </select>
        
        <input
          type="text"
          value={test.path}
          onChange={(e) => updateTest('path', e.target.value)}
          placeholder="/api/endpoint"
        />
      </div>
      
      {test.method !== 'GET' && (
        <div className="request-body">
          <label>Request Body</label>
          <textarea
            value={test.body || ''}
            onChange={(e) => updateTest('body', e.target.value)}
            placeholder='{"key": "value"}'
          />
        </div>
      )}
      
      <AssertionList
        assertions={test.assertions}
        onChange={(assertions) => updateTest('assertions', assertions)}
      />
    </div>
  );
};

export default TestCaseForm;