package semaphore_test

import (
	"github.com/andrewmolyuk/pixar/semaphore"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewSemaphore(t *testing.T) {
	//Act
	s := semaphore.NewSemaphore(5)

	//Assert
	assert.NotNil(t, s, "Semaphore should not be nil")
	assert.Equal(t, 5, s.Limit(), "Semaphore limit should be 5")
}

func TestSemaphore_Add(t *testing.T) {
	//Arrange
	s := semaphore.NewSemaphore(3)

	//Act
	s.Acquire()

	//Assert
	assert.Equal(t, 1, s.Count(), "Semaphore count should be 1")
	assert.Equal(t, 3, s.Limit(), "Semaphore limit should be 3")
}

func TestSemaphore_Add_Failed(t *testing.T) {
	//Arrange
	s := semaphore.NewSemaphore(1)

	//Act
	s.Acquire()
	go func() {
		time.Sleep(time.Nanosecond)
		s.Release()
	}()
	s.Acquire()

	//Assert
	assert.Equal(t, 1, s.Count(), "Semaphore count should be 1")
}

func TestSemaphore_Done(t *testing.T) {
	//Arrange
	s := semaphore.NewSemaphore(1)
	s.Acquire()

	//Act
	s.Release()

	//Assert
	assert.Equal(t, 0, s.Count(), "Semaphore count should be 0")
}

func TestSemaphore_Wait(t *testing.T) {
	//Arrange
	s := semaphore.NewSemaphore(1)

	//Act
	s.Acquire()
	s.Release()
	s.Wait()

	//Assert
	assert.Equal(t, 0, s.Count(), "Semaphore count should be 0")
}
