package pwd

import (
    "testing"
)


type MockDirectoryRetriever struct {
    MockDir string
    MockErr error
}

func (m MockDirectoryRetriever) Getwd() (string, error) {
    return m.MockDir, m.MockErr
}
func TestDefaultDirectoryRetriever(t *testing.T) {
    // This test uses the actual file system and should be used to ensure the real implementation is covered.
    retriever := DefaultDirectoryRetriever{}
    _, err := CurrentWorkingDirectory(retriever)
    if err != nil {
        t.Errorf("Failed to retrieve working directory using DefaultDirectoryRetriever: %s", err)
    }
}

func TestCurrentWorkingDirectory(t *testing.T) {
    cases := []struct {
        name       string
        mockDir    string
        mockErr    error
        expectDir  string
        expectErr  error
    }{
        {
            name:       "successful retrieval",
            mockDir:    "/home/user",
            expectDir:  "/home/user",
            expectErr:  nil,
        },
        {
            name:       "error retrieval",
            mockErr:    fmt.Errorf("error getting directory"),
            expectErr:  fmt.Errorf("error getting directory"),
        },
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            retriever := MockDirectoryRetriever{MockDir: tc.mockDir, MockErr: tc.mockErr}
            dir, err := CurrentWorkingDirectory(retriever)
            if dir != tc.expectDir || (err != nil && tc.expectErr != nil && err.Error() != tc.expectErr.Error()) {
                t.Errorf("Test %s failed: expected %v, %v got %v, %v", tc.name, tc.expectDir, tc.expectErr, dir, err)
            }
        })
    }
}

