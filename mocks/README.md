# Generated Mocks

This directory contains mocks generated using [mockery](https://github.com/vektra/mockery) for the interfaces defined in the `interfaces` package.

## Usage

### Basic Usage

```go
import (
    "testing"
    "github.com/pedropazello/url-redirect-service/mocks"
)

func TestSomething(t *testing.T) {
    // Create a new mock
    mockRepo := mocks.NewIRedirectsRepository(t)
    
    // Set expectations
    mockRepo.EXPECT().GetItem(ctx, "test-id").Return(expectedResult, nil)
    
    // Use the mock in your test
    // ... test code here
    
    // Mock expectations are automatically verified when test ends
}
```

### Advanced Usage with Matchers

```go
import (
    "github.com/stretchr/testify/mock"
)

func TestWithMatchers(t *testing.T) {
    mockRepo := mocks.NewIRedirectsRepository(t)
    
    // Use mock.Anything for flexible matching
    mockRepo.EXPECT().GetItem(mock.Anything, mock.AnythingOfType("string")).Return(result, nil)
    
    // Use mock.MatchedBy for custom matching logic
    mockRepo.EXPECT().GetItem(
        mock.Anything, 
        mock.MatchedBy(func(id string) bool {
            return strings.HasPrefix(id, "test-")
        })
    ).Return(result, nil)
}
```

## Regenerating Mocks

To regenerate the mocks after interface changes:

```bash
# Using Docker (recommended)
make mocks

# Or directly with Docker
docker run --rm -v $(pwd):/src -w /src vektra/mockery:v2.46.3

# Clean and regenerate
make regenerate-mocks
```

## Configuration

Mock generation is configured in `.mockery.yaml` in the project root.

## Notes

- All mocks are automatically generated - do not edit manually
- Mock expectations are automatically verified at test completion
- Use the `EXPECT()` pattern for type-safe expectation setting
- The constructor `NewXXX(t)` automatically handles cleanup and assertion
